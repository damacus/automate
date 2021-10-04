package reportmanager

import (
	"context"
	"fmt"
	"net"

	"github.com/chef/automate/api/interservice/report_manager"
	"github.com/chef/automate/components/report-manager-service/config"
	"github.com/chef/automate/components/report-manager-service/server"
	"github.com/chef/automate/lib/grpc/health"
	"github.com/chef/automate/lib/grpc/secureconn"
	"github.com/sirupsen/logrus"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Serve(conf config.ReportManager, connFactory *secureconn.Factory) error {
	ctx := context.Background()

	//get minio connection
	minioClient, err := getMinioConnection(ctx, conf)
	if err != nil {
		logrus.WithError(err).Fatal("Error in establishing a connection to minio")
		return err
	}
	logrus.Infof("connection established to minio server, endPoint:%s", minioClient.EndpointURL())

	return serveGrpc(ctx, conf, minioClient, connFactory)
}

func getMinioConnection(ctx context.Context, conf config.ReportManager) (*minio.Client, error) {

	//TODO:: Get the below details from configuration
	endpoint := "127.0.0.1:10197"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("error in creating a connection to minio server: %w", err)
	}

	//check if the default bucket is available or not
	isExist, err := minioClient.BucketExists(ctx, "default")
	if err != nil {
		return nil, fmt.Errorf("error in checking the default bucket existence in minio server:%w", err)
	}
	//create a default bucket if not available
	if !isExist {
		err := minioClient.MakeBucket(ctx, "default", minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("error in creating a default bucket to store the report data: %w", err)
		}
	}

	return minioClient, nil
}

func serveGrpc(ctx context.Context, conf config.ReportManager, minioClient *minio.Client, connFactory *secureconn.Factory) error {

	grpcBinding := fmt.Sprintf("%s:%d", conf.Service.Host, conf.Service.Port)
	lis, err := net.Listen("tcp", grpcBinding)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := connFactory.NewServer()

	//register health server for health status
	health.RegisterHealthServer(s, health.NewService())
	report_manager.RegisterReportManagerServiceServer(s, server.New(minioClient))

	logrus.Info("Starting GRPC server on " + grpcBinding)

	return s.Serve(lis)
}
