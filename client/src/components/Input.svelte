<script>
  export let value = "";
  export let size = "regular";

  const SIZE = {
    large: "text-lg px-3.5 py-2 rounded-lg",
    regular: "text-base px-3 py-1.5 rounded-md",
    small: "text-sm px-2.5 py-1 rounded-sm",
  };

  const commonStr = "inline-flex items-center ";
  const styleStr = "bg-white border";
  const sizeStr = SIZE[size ?? "regular"];
</script>

<input type="text" bind:value class={`${commonStr} ${styleStr} ${sizeStr}`} />
