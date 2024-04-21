<script>
  import Button from "@/components/Button.svelte";
  import Input from "@/components/Input.svelte";

  export let items = [];

  const addItem = () => {
    items = [...items, "New Section"];
  };

  const removeItem = (idx) => {
    items = [...items.slice(0, idx), ...items.slice(idx + 1)];
  };
</script>

<ul class="list-none flex flex-col gap-3 w-full">
  {#each items as _item, idx}
    <li class="flex gap-2">
      <Input bind:value={items[idx]} extraClass="flex-1" />
      <Button
        label="Remove"
        onClick={() => removeItem(idx)}
        disabled={items.length === 1}
      />
    </li>
  {/each}
</ul>
<Button label="Add" onClick={addItem} />
