import logo from './logo.svg';
import './App.css';
import Articles from './Components/Articles';
import AddArticle from './Components/AddArticle';
import "bootswatch/dist/sandstone/bootstrap.min.css";

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useRouteMatch,
  useParams
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Router>
        <div>
          <ul>
            <li>
              <Link to="/articles">Articles list</Link>
            </li>
            <li>
              <Link to="/add">Add Article</Link>
            </li>
          </ul>

          <Switch>
            <Route path="/articles">
              <Articles></Articles>
            </Route>
            <Route path="/add">
              <AddArticle></AddArticle>
            </Route>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
