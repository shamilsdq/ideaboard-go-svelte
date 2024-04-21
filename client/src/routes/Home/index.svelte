<script>
  import CreateSection from "./CreateSection.svelte";
  import JoinSection from "./JoinSection.svelte";
</script>

<div class="w-screen min-h-screen flex justify-center items-center">
  <main class="flex flex-col gap-10 max-w-4xl">
    <h2 class="text-5xl font-app">Idea Board for teams</h2>
    <div class="flex gap-7">
      <CreateSection />
      <JoinSection />
    </div>
  </main>
</div>
