## Anotações importantes projeto GO

##Camadas da aplicação
Domain -> Camada com as regras do negócios, somente a lógica do negócio é tratada nela
Framework -> Camada com as complecidades técnicas  da aplicação
Aplications -> Camada que utiliza o Domain e Framework para realizar as tarefas necessárias
## Passo a seguir
Iniciar o docker ' docker-compose up -d '
Utilitário 'go mod' gerenciador de pacotes go
iniciano $go mod init package-name
executando $go run file_name.go
Adicionar a credencial
export GOOGLE_APPLICATION_CREDENTIALS="/go/src/codeeducation-test-260423-7501a56d1b0e.json"
