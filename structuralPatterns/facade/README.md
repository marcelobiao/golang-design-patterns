# Facade Pattern


O Facade é um structural pattern que provê uma interface  simplificada para a utilização de um biblioteca, framework ou qualquer outra classe complexa

Imagine que você precisa fazer seu código funcionar com uma biblioteca sofisticada. Normalmente, você precisaria instanciar e configurar todos os objetos que essa lib utiliza e executar os métodos na ordem correta. Porém, isso acaba se tornando um problema, pois, dessa forma, as regras de negócio da sua aplicação ficam fortemente acopladas com os detalhes de implementação das classes da lib, dificultando a compreensão do código e sua manutenção.

Através do facade é possível acessar os recursos de uma lib através de uma interface simples, que oculta detalhes mais complexos de implementação. O Facade fornece apenas os recursos que os clientes realmente se importam, o que é bom por um lado, pois sua utilização acaba se tornando mais simples, porém, pode fornecer funcionalidades mais limitadas em comparação ao acesso direto às classes.

Uma boa analogia da vida real seria analisar o processo de uma compra online. Depois que adicionamos os produtos ao nosso carrinho de compras e realizamos o pagamento, esses produtos passam por toda uma cadeia logística envolvendo sua separação, embalagem e transporte, através de caminhoẽs e aviões, passando por diversos centros logísticos até chegar em nossa residência. Imagine o quão complexo seria para o usuário ter que se preocupar com todos esses detalhes ao realizar uma compra online.

Ao invés disso, as empresas oferecem para o cliente uma "Facede", que abstrai todos esses detalhes, ficando sob sua responsabilidade do usuário apenas escolher os produtos, informar seus dados e realizar o pagamento.


## Quando usar o Facade

- Use o Facade quando precisar fornecer uma interface limitada, porém direta para um sistema complexo.
- Use a Facade quando desejar estruturar um subsistema em camadas.

## Como implementar o Facade
- Verifique se é possível fornecer uma interface mais simples do que o próprio sistema já fornece.
- Declare e implemente sua interface. O Facade deve redirecionar as chamadas do código do cliente para as classes apropriadas do subsistema.
- O ideal, para tirar todo o proveito do padrão, é fazer com que o cliente se comunique apenas com o Facade. Assim, o cliente estará protegido contra qualquer alteração dos códigos do subsistema, desde que o Facade seja atualizado também.
