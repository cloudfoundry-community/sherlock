# Sherlock

What is Sherlock? Sherlock is used for auditing purposing. It can be used to notify when "break glass" account is used such as admin. Sherlock looks for users login in the  [uaa](https://github.com/cloudfoundry/uaa) logs. It uses [audit](https://github.com/cloudfoundry/uaa/blob/master/docs/UAA-Audit.rst) for `UserAuthentication` audits in log.

## Usage


```
sherlock -log uaa.log -time 60s -users admin
```

```
Options:
-log: Location of uaa log
-time: Amount of time to back in log to look for users
-users: Comma seperate list of users to search for
```

Sherlock returns exit 1 on finding a user tried to log in. This can be hook up to things like nagios or sensu for alerting.
