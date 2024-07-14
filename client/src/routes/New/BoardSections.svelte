<script>
  import { PlusSolid, TrashCanRegular } from "svelte-awesome-icons";
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
        label=""
        style="transparent"
        onClick={() => removeItem(idx)}
        disabled={items.length === 1}
        iconClass={TrashCanRegular}
      />
    </li>
  {/each}
</ul>
<Button
  label="Add new section"
  onClick={addItem}
  iconClass={PlusSolid}
  style="secondary"
/>
