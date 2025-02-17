# encoding: utf-8
# copyright: 2019, Chef Software, Inc.
# license: All rights reserved

title 'Chef Automate Backend Deploy Smoke Tests'

%w(
  chef/automate-backend-postgresql
  chef/automate-backend-haproxy
  chef/automate-backend-journalbeat
  chef/automate-backend-metricbeat
  chef/automate-backend-pgleaderchk
).each do |svc|
  describe hab_svc(svc) do
    it { should be_installed }
    it { should be_up }
  end
end

%w(
  postgresql
  haproxy
  pgleaderchk
).each do |svc|
  describe command("HAB_LICENSE=accept-no-persist /hab/svc/automate-backend-#{svc}/hooks/health-check") do
    its('exit_status') { should eq 0 }
  end
end

log_file = command('ls -t /hab/svc/automate-backend-postgresql/var/pg_log/').stdout.lines.first.chomp

describe command("tail -10 /hab/svc/automate-backend-postgresql/var/pg_log/#{log_file} | grep \"FATAL\\|PANIC\"") do
  its('stdout') { should match /^\s*$/ }
  its('exit_status') { should eq 1 } # exit 1 means grep did found no pattern matches
end

describe x509_certificate('/hab/svc/automate-backend-postgresql/config/server.crt') do
  its('validity_in_days') { should be > 30 }
end
