## 2.0.5
 - remove implementation specific test in specs

## 2.0.3
 - fixed specs for time handling

## 2.0.2
 - added comments and basic specs

## 2.0.0
 - Plugins were updated to follow the new shutdown semantic, this mainly allows Logstash to instruct input plugins to terminate gracefully,
   instead of using Thread.raise on the plugins' threads. Ref: https://github.com/elastic/logstash/pull/3895
 - Dependency on logstash-core update to 2.0