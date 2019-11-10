# Dicas de GitHub

## VOLTANDO PRO MASTER

caso modificou algo e deseja voltar pro master

### retornado todos os arquvio

        git checkout .

### retornado 1 arquivo 

        git checkout /caminho/file.name 

### voltando para branch do master

        git checout origin master
## TRABALHANDO COM BRANCH

caso tenha alterado alguma arquivo pode realizar o push em uma branch

#### CRIA BRANCH

    git checkout -b nome-da-sua-branch


#### POSTANDO NA BRANCH

* ADD LINHAS, No caso de não precisar subir novos arquivos.

        git add -p

* ADD ARQUIVO ESPECIFICO

        git add ./caminho/arquivo.nome

* ADD TODA UMA PASTA (NÃO RECOMENDADO)

        git add  ./pasta/

### MERGE, POSTANDO A BRANCH NO MASTER

sigua para:
https://github.com/alypher/lip/branches

botão:NEW PULL REQUEST -> CREATE PULL REQUEST