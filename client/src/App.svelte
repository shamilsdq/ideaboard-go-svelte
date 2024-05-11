<script>
  import { Router, Route } from "svelte-routing";
  import { Home } from "@/routes/Home";
  import { New } from "@/routes/New";
  import { Board } from "./routes/Board";
</script>

<Router>
  <Route path="/" component={Home} />
  <Route path="/board/new" component={New} />
  <Route path="/board/:boardId" component={Board} />
  <Route>Not found</Route>
</Router>
