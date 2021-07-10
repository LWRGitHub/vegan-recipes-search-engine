import React from 'react';
import {
  HashRouter,
  Route,
  Switch
} from 'react-router-dom';

// Components
import About from './About';
import Home from './Home';
import Footer from './Footer';

const App = () => (
  <HashRouter>
    <div>
      <Switch>
        <Route path="/home" component={Home} />
        <Route path="/about" component={About} />
        <Route component={Home} />
      </Switch>
      <Footer />
    </div>
  </HashRouter>
)

export default App;
