<script>
  import { createEventDispatcher } from "svelte";
  import { CirclePlusSolid } from "svelte-awesome-icons";
  import { boardStore } from "@/stores/board";
  import Post from "./Post.svelte";

  export let section;

  const dispatch = createEventDispatcher();

  $: posts = $boardStore.posts?.filter((post) => post.sectionId === section.id);

  const addPost = () => dispatch("add-post");
</script>

<section class="flex-1 min-w-80 flex flex-col bg-gray-100 rounded-sm border">
  <div class="w-full flex gap-5 p-3.5 items-start">
    <h3 class="flex-1 font-app font-medium leading-6 break-words">
      {section.title}
    </h3>
    <div title="Add new post" class="rounded-full">
      <CirclePlusSolid
        withEvents
        size="20"
        role="button"
        class="rounded-full duration-100 hover:scale-110 focus:outline-2 focus:outline-offset-2"
        ariaLabel="Hello World"
        on:click={addPost}
      />
    </div>
  </div>
  <div class="masonry p-3 gap-2 overflow-y-auto">
    {#each posts as post}
      <Post data={post} />
    {/each}
  </div>
</section>

<style>
  .masonry {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  }
</style>
