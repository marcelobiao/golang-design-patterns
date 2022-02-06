# Singleton Pattern


O Singleton é um Creational Pattern responsável por garantir que uma classe tenha apenas uma instância, enquanto fornece um ponto de acesso global para a mesma.

Sempre que chamamos o construtor de uma classe, o mesmo fica responsável por gerar uma nova instância e retornar para quem realizou a chamada. Assim, a cada nova chamada do construtor, temos uma nova instância da classe em nosso sistema.

Com o Singleton isso muda bastante, pois na primeira vez que o construtor é invocado, uma nova instância é criada, salva em um cache e retornada para quem realizou a chamada. Porém, a partir da segunda chamada do construtor, o Singleton ao invés de gerar uma nova instância, retorna o cache da instância criada anteriormente, fazendo com que exista apenas uma instância da classe, que será compartilhada por todos.

O motivo mais comum para a adoção desse pattern é controlar o acesso a um recurso compartilhado, como um arquivo ou base de dados.

## Implementação:
- 1- Torne privado o construtor padrão da classe. Isso evita que novas instâncias dessa classe sejam geradas no código.
- 2- Crie um método estático que atue como um construtor. Esse método por baixo dos panos irá chamar o construtor privado para criar um objeto e salvar em um cache. Todas as chamadas a seguir, irão retornar o objeto armazenado em cache.

## Vantagens:
- Você terá certeza de que existe apenas uma instância de sua classe.
- Você adquire um ponto de acesso global para sua instância.
- O objeto Singleton é inicializado apenas quando é solicitado pela primeira vez.
