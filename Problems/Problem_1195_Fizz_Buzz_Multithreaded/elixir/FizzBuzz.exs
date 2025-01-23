defmodule FizzBuzz do
  def start(n) do
    # Crear procesos para cada función
    fizz_pid = spawn(FizzBuzz, :fizz, [])
    buzz_pid = spawn(FizzBuzz, :buzz, [])
    fizzbuzz_pid = spawn(FizzBuzz, :fizzbuzz, [])
    number_pid = spawn(FizzBuzz, :number, [])

    # Crear el proceso principal (coordinar) que llevará el control
    main_pid = spawn(FizzBuzz, :main, [1, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid])

    # Iniciar el ciclo de ejecución
    send(main_pid, :start)
  end

  def main(i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid) do
    receive do
      :start ->
        # Iniciar el ciclo con el número 1
        run(i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)
    end
  end

  def run(i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid) do
    cond do
      i <= n ->
        cond do
          rem(i, 3) == 0 and rem(i, 5) == 0 ->
            send(fizzbuzz_pid, {:fizzbuzz, i})
            receive_message(:fizzbuzz, i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)

          rem(i, 3) == 0 ->
            send(fizz_pid, {:fizz, i})
            receive_message(:fizz, i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)

          rem(i, 5) == 0 ->
            send(buzz_pid, {:buzz, i})
            receive_message(:buzz, i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)

          true ->
            send(number_pid, {:number, i})
            receive_message(:number, i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)
        end

      i > n ->
        # Notificar a los procesos que terminen
        Enum.each([fizz_pid, buzz_pid, fizzbuzz_pid, number_pid], fn pid ->
          send(pid, :done)
        end)
    end
  end

  def receive_message(process, i, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid) do
    receive do
      :done ->
        # Después de que un proceso termine su tarea, pasa el control al siguiente
        run(i + 1, n, fizz_pid, buzz_pid, fizzbuzz_pid, number_pid)
    end
  end

  # Funciones para cada uno de los procesos
  def fizz do
    receive do
      {:fizz, _i} ->
        IO.puts("fizz")
        send(self(), :done)
        fizz()

      :done ->
        :ok
    end
  end

  def buzz do
    receive do
      {:buzz, _i} ->
        IO.puts("buzz")
        send(self(), :done)
        buzz()

      :done ->
        :ok
    end
  end

  def fizzbuzz do
    receive do
      {:fizzbuzz, _i} ->
        IO.puts("fizzbuzz")
        send(self(), :done)
        fizzbuzz()

      :done ->
        :ok
    end
  end

  def number do
    receive do
      {:number, i} ->
        IO.puts(i)
        send(self(), :done)
        number()

      :done ->
        :ok
    end
  end
end

# Ejecutar FizzBuzz para un valor n
FizzBuzz.start(15)
