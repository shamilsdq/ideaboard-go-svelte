<script>
  import { Router, Route } from "svelte-routing";
  import { Home } from "@/routes/Home";
  import { New } from "@/routes/New";
</script>

<Router>
  <Route path="/" component={Home} />
  <Route path="/board/new" component={New} />
  <Route>Not found</Route>
</Router>
