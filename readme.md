projeto investlab 2.0

para  vc simular comprar e venda de ação



precisa instalar na sua maquina,

go, 
nodejs, 
yarn, 
ansible, 
( para usar o ansible)
python 
caso for mac vai precisar do brew instalado

para fazer o build precisa alterar algumas variaveis do ansible ainda não está automarica

./playbook/group_vars/all.yml

mude para seu diretorio caso queira saber execute o comando 'pwd' no terminal após a alteração
para fazer o build do projeto basta executar o build.sh na pasta raiz do projeto
para subir o servidor vai na pasta core/src/server e execute o comando 'go run main.go'

## Install go

    sudo apt install go

## Install node

    sudo apt install nodejs

## Install yarn

    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

    sudo apt-get update && sudo apt-get install yarn
    sudo apt-get install --no-install-recommends yarn

## Build PROJEC


    ./build.sh

#### build front-end

    cd ./ui-browser/
    npm run build

#### run server

    cd core/src/server
    go run main.go

open: http://localhost:8080/

#### run front-end mode developer

    cd ./ui-browser/
    npm run serve


open: http://localhost:8084/