<script>
  import { Router, Route } from "svelte-routing";
  import { Home } from "@/routes/Home";
</script>

<Router>
  <Route path="/" component={Home} />
  <Route>Not found</Route>
</Router>
