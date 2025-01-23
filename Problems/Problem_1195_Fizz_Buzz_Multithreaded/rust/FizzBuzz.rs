use std::sync::{Arc, Mutex}; // Arc y Mutex para sincronización entre hilos
use std::thread;

// Estructura principal para FizzBuzz
struct FizzBuzz {
    n: usize, // Número máximo hasta el cual se ejecutará FizzBuzz
}

impl FizzBuzz {
    // Constructor para inicializar la estructura
    fn new(n: usize) -> Self {
        Self { n }
    }

    /// Método genérico para manejar una acción basada en una condición
    ///
    /// - `condition`: Un cierre (closure) que toma un usize y retorna un bool.
    /// - `action`: Un cierre (closure) que ejecuta una acción basada en la condición.
    fn run(
        &self,
        condition: impl Fn(usize) -> bool + Send + Sync + 'static,
        action: Arc<dyn Fn() + Send + Sync>, // Usamos Arc para permitir compartir entre hilos
    ) {
        let n = self.n; // Capturamos el valor máximo para el bucle
        thread::spawn(move || {
            for i in 1..=n {
                if condition(i) {
                    action(); // Ejecutamos la acción si la condición se cumple
                }
            }
        })
        .join()
        .expect("El hilo no pudo completarse correctamente.");
    }

    /// Lógica para "Fizz" (números divisibles por 3 pero no por 5)
    fn fizz(&self, print_fizz: Arc<dyn Fn() + Send + Sync>) {
        self.run(
            |x| x % 3 == 0 && x % 5 != 0, // Condición: divisible por 3 pero no por 5
            print_fizz,
        );
    }

    /// Lógica para "Buzz" (números divisibles por 5 pero no por 3)
    fn buzz(&self, print_buzz: Arc<dyn Fn() + Send + Sync>) {
        self.run(
            |x| x % 5 == 0 && x % 3 != 0, // Condición: divisible por 5 pero no por 3
            print_buzz,
        );
    }

    /// Lógica para "FizzBuzz" (números divisibles por 3 y 5)
    fn fizzbuzz(&self, print_fizzbuzz: Arc<dyn Fn() + Send + Sync>) {
        self.run(
            |x| x % 3 == 0 && x % 5 == 0, // Condición: divisible por 3 y 5
            print_fizzbuzz,
        );
    }

    /// Lógica para números que no son divisibles ni por 3 ni por 5
    fn number(&self, print_number: Arc<dyn Fn(usize) + Send + Sync>) {
        let n = self.n;
        thread::spawn(move || {
            for i in 1..=n {
                if i % 3 != 0 && i % 5 != 0 {
                    print_number(i); // Ejecutamos la función para imprimir el número
                }
            }
        })
        .join()
        .expect("El hilo no pudo completarse correctamente.");
    }
}

fn main() {
    // Número máximo hasta el cual se ejecutará FizzBuzz
    let n = 15;

    // Crear una nueva instancia de FizzBuzz
    let fizz_buzz = FizzBuzz::new(n);

    // Definir las acciones (usando closures envueltos en Arc)
    let print_fizz = Arc::new(|| println!("Fizz"));
    let print_buzz = Arc::new(|| println!("Buzz"));
    let print_fizzbuzz = Arc::new(|| println!("FizzBuzz"));
    let print_number = Arc::new(|x: usize| println!("{}", x));

    // Ejecutar cada lógica en threads separados
    fizz_buzz.fizz(print_fizz.clone()); // Lógica para "Fizz"
    fizz_buzz.buzz(print_buzz.clone()); // Lógica para "Buzz"
    fizz_buzz.fizzbuzz(print_fizzbuzz.clone()); // Lógica para "FizzBuzz"
    fizz_buzz.number(print_number.clone()); // Lógica para números restantes
}
