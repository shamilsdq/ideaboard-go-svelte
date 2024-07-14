<script>
  import { createEventDispatcher, onMount } from "svelte";
  import { blur } from "svelte/transition";
  import Button from "@/components/Button.svelte";

  export let data;

  const dispatch = createEventDispatcher();

  let inputElement;
  let content = "";

  const handleCancel = () => {
    dispatch("close");
  };

  const handleSave = () => {
    const payload = { sectionId: data.id, content };
    dispatch("save", payload);
    handleCancel();
  };

  const handleKeydown = (event) => {
    if (event.key === "Escape") handleCancel();
  };

  onMount(() => {
    inputElement?.focus();
  });
</script>

<div
  class="fixed inset-0 flex items-center justify-center bg-black/65 backdrop-blur-sm"
  role="presentation"
  tabindex="-1"
  on:keydown={handleKeydown}
  on:click={handleCancel}
  transition:blur={{ duration: 200 }}
>
  <div
    class="w-96 bg-white p-5 rounded-lg shadow-xl space-y-5"
    role="presentation"
    on:click|stopPropagation
  >
    <header class="font-app text-xs flex items-center gap-2">
      <span class="px-2 py-1 bg-gray-100 border rounded-sm">{data.title}</span>
      <span>&gt;</span>
      <span>New post</span>
    </header>
    <div
      contenteditable="true"
      role="textbox"
      placeholder="Enter text"
      class="w-full max-h-40 min-h-24 overflow-auto resize-none p-2 border rounded focus:outline-none focus:border-gray-400"
      bind:innerText={content}
      bind:this={inputElement}
    ></div>
    <footer class="flex gap-2 justify-end">
      <Button label="Cancel" onClick={handleCancel} />
      <Button label="Save" onClick={handleSave} />
    </footer>
  </div>
</div>
