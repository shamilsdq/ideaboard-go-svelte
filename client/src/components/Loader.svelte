<script>
  export let style = "primary";
  export let size = "regular";

  const STYLE = {
    primary: "border-gray-300 border-t-blue-600",
  };

  const SIZE = {
    large: "h-20 w-20 border-8",
    regular: "h-14 w-14 border-8",
    small: "h-8 w-8 border-4",
  };

  const commonStr = "rounded-full animate-spin";
  const sizeStr = SIZE[size ?? "regular"];
  const styleStr = STYLE[style ?? "primary"];
</script>

<div class={`${commonStr} ${styleStr} ${sizeStr}`} />
