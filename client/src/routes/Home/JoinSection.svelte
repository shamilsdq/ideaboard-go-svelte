<script>
  import Button from "@/components/Button.svelte";
  import Input from "@/components/Input.svelte";
</script>

<section
  class="border flex-1 flex flex-col justify-between items-start gap-5 p-7 bg-gray-100 rounded-md"
>
  <p class="leading-relaxed">
    Be an active team player. Join your team in their brainstorming session and
    pitch in your ideas and opinions.
  </p>
  <div class="flex gap-2">
    <Input />
    <Button label="Join Board" />
  </div>
</section>
