<script>
  import boardsService from "@/services/boards";
  import { onMount } from "svelte";

  export let boardId;

  let socket = undefined;

  onMount(async () => {
    socket = await boardsService.connectSocket(boardId);
  });
</script>

<div>Board</div>
