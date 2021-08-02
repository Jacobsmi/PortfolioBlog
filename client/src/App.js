import { BrowserRouter, Switch, Route } from "react-router-dom";
import Landing from "./landing/landing";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Switch>
          <Route exact path="/">
            <Landing />
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
