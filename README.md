# loginWithTokenGo

Função lambda AWS implementada em Go de validação de login.

Utiliza uma Autorização passada como Header de uma requisição HTTP para chamada de uma API externa de autenticação.

Configurações AWS:
1) Deve ser configurada a Variável de Ambiente da AWS "LoginUnicoURL" com o endereço da API externa de autenticação.
2) Deve ser liberado acesso da função Lambda para o recurso Amazon CloudWatch Logs.
3) Configurar uma API Gateway como gatilho para a função Lambda.
4) Tempo limite foi configurado para 3s e a memória em 128MB.
