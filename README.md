## Implementação do Algortimo H1

    Algoritmo heurístico (H1): Um algoritmo heurístico baseado no endereço de entrada da transação, heurística de propriedade de entrada comum. De acordo
    com o protocolo do sistema Bitcoin, se você quiser usar o bitcoin de um endereço, a chave privada desse o endereço deve ser fornecido,o que significa
    que o usuário do endereço deve assinar para a transação. Consequentemente, quando múltiplos endereços são usados como a entrada de uma transação em 
    conjunto, acreditamos que todos os endereços de entrada da transação podem ser agrupados em um grupo de endereços. Em outras palavras, todos os endereços
    de entrada são controlados pela mesma entidade de transação. A precisão do clustering de H1 pode chegar a 100% sem considerar o fato de que os usuários 
    usam serviços de mistura para evitar o clustering análise intencionalmente.

### Building Database

Salva os elementos que serão analisados para extrair ou tratar as transações.

### Process Transactions

Salva as transações

### Create Cluster

Cria os Cluster com as transações

### Process Cluster

Utiliza o Algoritmo H1 nos Clusters salvos
