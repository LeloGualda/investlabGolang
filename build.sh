./tools/create-local-var.sh

ansible-playbook -K  -vvvv -i ./build/playbook/hosts ./build/playbook/provisioning.yml