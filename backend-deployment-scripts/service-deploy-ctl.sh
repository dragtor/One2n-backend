#!bin/bash
#restart : stop supervior & loadenv & start service running on location
#start : load env & start service
#stop : stop supervisor
#deploy : pull latest code of repository , generate artifact & place it to desire location & start service

OPTION=$1

case $OPTION in
##############################  deploy #################
    help)   
          echo "commands to run :"
          echo "1. deploy [first time deployment]"
          echo "2. update [pull lataest code & deploy]"
          echo "3. status [check service status]"
        ;;
    deploy)
            git clone https://github.com/dragtor/One2n-backend.git temp
            cd $HOME/temp/backend
            make build
            #copy artifact & supervisorfile & .env file
            mkdir $HOME/s3ls-webservice
            cp s3ls $HOME/s3ls-webservice
            cp supervisord.conf $HOME/s3ls-webservice
            mkdir $HOME/s3ls-webservice/logs/
            cd $HOME
            rm -rf temp
            cd $HOME/s3ls-webservice
            supervisord -c supervisord.conf
            cd $HOME
        ;;
    update)
            git clone https://github.com/dragtor/One2n-backend.git temp
            cd $HOME/temp/backend
            make build
            cd $HOME/s3ls-webservice
            supervisorctl -c supervisord.conf stop all
            cd $HOME/temp/backend
            cp s3ls $HOME/s3ls-webservice
            cp supervisord.conf $HOME/s3ls-webservice
            cd $HOME
            rm -rf temp
            cd $HOME/s3ls-webservice
            supervisorctl -c supervisord.conf start all
            cd $HOME
        ;;
    status)
            cd $HOME/s3ls-webservice
            supervisorctl -c supervisord.conf status 
            cd $HOME
esac




