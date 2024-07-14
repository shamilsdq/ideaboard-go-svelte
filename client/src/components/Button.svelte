<script>
  export let label = "Button";
  export let size = "regular";
  export let style = "primary";
  export let onClick = undefined;
  export let disabled = false;
  export let iconClass = undefined;

  const STYLE = {
    primary: "text-white bg-gray-700 hover:bg-black disabled:bg-gray-400",
    secondary:
      "text-black bg-gray-300 hover:bg-gray-400 disabled:text-gray-500",
    transparent: "text-black hover:bg-gray-300 disabled:text-gray-400",
  };

  const SIZE = {
    large: "text-lg px-5 py-2 rounded-lg",
    regular: "text-base px-4 py-1.5 rounded-md",
    small: "text-sm px-3 py-1 rounded-sm",
  };

  const commonStr =
    "inline-flex gap-2 items-center duration-300 disabled:cursor-not-allowed";
  const styleStr = STYLE[style ?? "primary"];
  const sizeStr = SIZE[size ?? "regular"];
</script>

<button
  on:click={onClick}
  {disabled}
  class={`${commonStr} ${styleStr} ${sizeStr}`}
>
  {label}
  {#if iconClass != null}
    <svelte:component this={iconClass} size="14" />
  {/if}
</button>
