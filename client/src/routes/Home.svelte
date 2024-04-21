<script>
  import Button from "@/components/Button.svelte";
  import Input from "@/components/Input.svelte";
</script>

<div class="w-screen min-h-screen flex justify-center items-center">
  <main class="flex flex-col gap-10 max-w-4xl">
    <h2 class="text-5xl font-app">Idea Board for teams</h2>

    <div class="flex gap-7">
      <section
        class="border flex-1 flex flex-col justify-between items-start gap-5 p-7 bg-gray-100 rounded-md"
      >
        <p class="leading-relaxed">
          Create a custom board for your team. Share the link with your team and
          start your digital brainstorming session.
        </p>
        <Button label="Create New Board" />
      </section>
      <section
        class="border flex-1 flex flex-col justify-between items-start gap-5 p-7 bg-gray-100 rounded-md"
      >
        <p class="leading-relaxed">
          Be an active team player. Join your team in their brainstorming
          session and pitch in your ideas and opinions.
        </p>
        <div class="flex gap-2">
          <Input />
          <Button label="Join Board" />
        </div>
      </section>
    </div>
  </main>
</div>
