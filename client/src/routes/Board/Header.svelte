<script>
  import { boardStore } from "@/stores/board";
</script>

<header class="w-full flex justify-between items-center gap-3">
  <h1 class="text-2xl font-app">{$boardStore?.title}</h1>
  <span class="px-2.5 py-1.5 bg-gray-100 border text-sm leading-2 rounded">
    Members: {$boardStore?.memberCount}
  </span>
</header>
