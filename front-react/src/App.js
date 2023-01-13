import './css/app.css';
import 'antd/dist/reset.css';
import React, {Component} from 'react';
import Home from './Home';
class App extends Component {
  render() {
    return <div className="App">
      <Home></Home>
    </div>
  }
}

export default App;
