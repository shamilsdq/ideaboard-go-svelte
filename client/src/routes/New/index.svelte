<script>
  import { navigate } from "svelte-routing";

  import boardsService from "@/services/boards";
  import Button from "@/components/Button.svelte";
  import Input from "@/components/Input.svelte";
  import BoardSections from "./BoardSections.svelte";

  let title = "";
  let sections = ["Section"];

  const onSubmit = async () => {
    const payload = { title, sections };
    try {
      const { boardId } = await boardsService.create(payload);
      if (boardId) navigate(`/board/${boardId}`);
    } catch (err) {
      console.error(err);
    }
  };
</script>

<div class="w-screen min-h-screen flex justify-center items-center py-16">
  <div class="flex flex-col gap-5 max-w-4xl">
    <section class="mb-8">
      <h1 class="text-5xl font-app">New Board</h1>
    </section>

    <section class="flex flex-col gap-3 p-8 rounded border bg-gray-100">
      <h3 class="text-xl">Board Title</h3>
      <Input extraClass="min-w-96" bind:value={title} />
    </section>

    <section
      class="flex flex-col gap-3 items-start p-8 rounded border bg-gray-100"
    >
      <h3 class="text-xl">Board Sections</h3>
      <BoardSections bind:items={sections} />
    </section>

    <section class="mt-2">
      <Button
        label="Create Room"
        onClick={onSubmit}
        disabled={title === "" || sections.length === 0}
        size="large"
      />
    </section>
  </div>
</div>
