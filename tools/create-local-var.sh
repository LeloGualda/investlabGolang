#create file
patch_var_local="./playbook/group_vars/local.yml"

LOCATION=$(pwd)
# create vars
echo "---" > $patch_var_local
echo "dir_local: $LOCATION" >> $patch_var_local
#echo "mydir_server: $LOCATION/core/server" >> $patch_var_local

# # create get diretory

# S=" "
# L1=$(cat  $patch_var_local | head -1 | tail -1)
# L21=$(cat $patch_var_local | head -2 | tail -1 | awk {'print $1'})
# L22=$(cat $patch_var_local | head -2 | tail -1 | awk {'print $2'})
# L31=$(cat $patch_var_local | head -3 | tail -1 | awk {'print $1'})
# L32=$(cat $patch_var_local | head -3 | tail -1 | awk {'print $2'})

# S1=$L1
# S2=$L21$S$LOCATION$L22
# S3=$L31$S$LOCATION$L32

# # set diretory
# echo $S1 > $patch_var_local
# echo $S2 >> $patch_var_local
# echo $S3 >> $patch_var_local