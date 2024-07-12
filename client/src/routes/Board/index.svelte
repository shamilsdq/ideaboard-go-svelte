<script>
  import { onMount } from "svelte";
  import boardsService from "@/services/boards";
  import Loader from "@/components/Loader.svelte";
  import { boardStore } from "@/stores/board";
  import Header from "./Header.svelte";
  import Section from "./Section.svelte";

  export let boardId;

  let socket = undefined;
  let isLoading = true;

  const onSocketOpened = (openedSocket) => {
    socket = openedSocket;
    isLoading = false;
  };

  const onSocketMessage = (message) => {
    const { errors, code, content } = JSON.parse(message.data);
    if (errors) console.error(errors);
    else if (code && content) boardStore.handle(code, content);
    else console.error(`Unexpected response: ${message.data}`);
  };

  const onSocketClosed = () => (socket = undefined);
  const onSocketError = (error) => console.error(error);

  onMount(() => {
    socket = boardsService.connectSocket(boardId);
    socket.onopen = onSocketOpened;
    socket.onclose = onSocketClosed;
    socket.onerror = onSocketError;
    socket.onmessage = onSocketMessage;
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
        <Section {section} />
      {/each}
    </div>
  {/if}
</div>
