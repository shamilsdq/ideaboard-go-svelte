<script>
  export let label = "Button";
  export let size = "regular";
  export let style = "primary";
  export let onClick = undefined;
  export let disabled = false;

  const STYLE = {
    primary: "text-white bg-blue-700 hover:bg-blue-800",
  };

  const SIZE = {
    large: "text-lg px-5 py-2 rounded-lg",
    regular: "text-base px-4 py-1.5 rounded-md",
    small: "text-sm px-3 py-1 rounded-sm",
  };

  const commonStr = "inline-flex items-center";
  const styleStr = STYLE[style ?? "primary"];
  const sizeStr = SIZE[size ?? "regular"];
</script>

<button
  on:click={onClick}
  {disabled}
  class={`${commonStr} ${styleStr} ${sizeStr}`}
>
  {label}
</button>
