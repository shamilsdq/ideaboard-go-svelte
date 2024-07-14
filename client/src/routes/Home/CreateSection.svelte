<script>
  import { navigate } from "svelte-routing";
  import { PlusSolid } from "svelte-awesome-icons";
  import Button from "@/components/Button.svelte";
</script>

<section
  class="border flex-1 flex flex-col justify-between items-start gap-5 p-7 bg-gray-100 rounded-md"
>
  <p class="leading-relaxed">
    Create a custom board for your team. Share the link with your team and start
    your digital brainstorming session.
  </p>
  <Button
    label="Create New Board"
    onClick={() => navigate("/board/new")}
    iconClass={PlusSolid}
  />
</section>
