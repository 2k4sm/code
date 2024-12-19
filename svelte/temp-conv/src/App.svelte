<script>
  class Temperature {
    #c = $state(0);
    #f = $state(0);
    
    get c() {
        return this.#c;
    }

    set c(c) {
        this.#c = c;
        this.#f = parseFloat((c * 9 / 5 + 32).toFixed(1));
    }

    get f() {
        return this.#f;
    }
    
    set f(f) {
        this.#f = f;
        this.#c = parseFloat(((f - 32) * 5 / 9).toFixed(1));
    }
  }
  let temperature = new Temperature();
</script>

<main class="flex flex-col bg-slate-500 min-h-fit h-[100vh]">
   <span class="text-center font-sans text-4xl text-purple-950 p-10">Temperature Converter</span>
   <div class=" flex justify-center align-center m-10 p-10 gap-10">
        <div class="flex flex-col gap-5">
          <span class="text-center font-sans text-2xl text-purple-950">Celsius</span>
          <input type="number" placeholder="Temperature in C" min=0 class="input input-bordered w-full max-w-xs" bind:value={temperature.c}/>
        </div>
        <div class="flex flex-col gap-5">
            <span class="text-center font-sans text-2xl text-purple-950">Fahrenheit</span>
            <input type="number" placeholder="Temperature in F" min=0 class="input input-bordered w-full max-w-xs" bind:value={temperature.f}/>
        </div>
   </div>
</main>