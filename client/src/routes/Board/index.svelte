<script>
  import { onDestroy, onMount } from "svelte";
  import boardsService from "@/services/boards";
  import postsService from "@/services/post";
  import Loader from "@/components/Loader.svelte";
  import { boardStore } from "@/stores/board";
  import Header from "./Header.svelte";
  import Section from "./Section.svelte";
  import FormModal from "./FormModal.svelte";

  export let boardId;

  let socket = undefined;
  let isLoading = true;
  let formModalData = null;

  const onMessage = (data) => {
    const { code, content } = data;
    if (code && content) boardStore.handle(code, content);
    else console.error(`Unexpected response: ${data}`);
  };

  const addPost = ({ detail }) => postsService.addPost(socket, detail);

  onMount(() => {
    boardsService.connectSocket(boardId, {
      loadingSetter: (state) => (isLoading = !!state),
      instanceSetter: (instance) => (socket = instance),
      messageHandler: onMessage,
    });
  });

  onDestroy(() => {
    socket?.close?.();
  });
</script>

<div
  class="w-full h-screen flex flex-col items-center justify-center p-5 gap-5"
>
  {#if isLoading}
    <Loader size="large" />
  {:else if $boardStore !== null}
    <Header />
    <div class="w-full flex-1 flex gap-3 overflow-y-hidden overflow-x-auto">
      {#each $boardStore?.sections as section}
        <Section {section} on:add-post={() => (formModalData = section)} />
      {/each}
    </div>

    {#if formModalData}
      <FormModal
        data={formModalData}
        on:close={() => (formModalData = null)}
        on:save={addPost}
      />
    {/if}
  {/if}
</div>
