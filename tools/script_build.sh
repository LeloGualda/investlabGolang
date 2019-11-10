#!/bin/bash

LOCATION=$(pwd)
L1=$(cat playbook/group_vars/all.yml | head -1 | tail -1)
L21=$(cat playbook/group_vars/all.yml | head -2 | tail -1 | awk {'print $1'})
L22=$(cat playbook/group_vars/all.yml | head -2 | tail -1 | awk {'print $2'})
L31=$(cat playbook/group_vars/all.yml | head -3 | tail -1 | awk {'print $1'})
L32=$(cat playbook/group_vars/all.yml | head -3 | tail -1 | awk {'print $2'})

#cat playbook/group_vars/all.yml | head -2 | tail -1 | awk {'print $1 " pwd"$2'} >> build.sh
#cat playbook/group_vars/all.yml | head -3 | tail -1 | awk {'print $1 " pwd"$2'} >> build.sh
#echo '' >> build.sh 
#echo 'ansible-playbook -K ./playbook/build-development.yml' >> build.sh

S1=$L1
S2=$L21$LOCATION$L22
S3=$L31$LOCATION$L32

#echo $S1 >> build.sh
#echo $S2 >> build.sh
#echo $S3 >> build.sh
#echo '' >> build.sh 
#echo 'ansible-playbook -K ./playbook/build-development.yml' >> build.sh


echo $S1 > playbook/group_vars/all.yml
echo $S2 >> playbook/group_vars/all.yml
echo $S3 >> playbook/group_vars/all.yml