#!/bin/bash
rabbitmqctl wait --timeout 60 /var/lib/rabbitmq/mnesia/rabbit\@$(hostname).pid
rabbitmqctl add_user appuser appuserpassword
rabbitmqctl set_permissions -p / appuser ".*" ".*" ".*"
