import React, { Component } from 'react';
import SideBar from './components/SideBar';
import Discover from './containers/Discover';
import NodeInfo from './containers/NodeInfo';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.menu = [
      {
        name: 'Node Information',
        component: <NodeInfo />
      },
      {
        'name': 'Discover',
        component: <Discover />
      }
    ];
    this.state = {
      title: this.menu[0].name,
      selectedItem: 0,
      toggle: false,
    }
  }
  collapse(e) {
    this.setState({
      toggle: !this.state.toggle,
    })
  }
  menuSelectedHandler(index) {
    this.setState({
      title: this.menu[index].name,
      selectedItem: index,
    })
  }
  render() {
    return (
      <div className="wrapper">
        <SideBar
          toggle={this.state.toggle}
          title="Gocho"
          menu={this.menu}
          onMenuSelected={this.menuSelectedHandler.bind(this)} />
        <div className="content-wrapper">
          <nav className="nav">
            <b>{this.state.title}</b>
            <button
              type="button"
              className="btn btn-info navbar-btn"
              onClick={this.collapse.bind(this)}
            >
              &lt;
            </button>
          </nav>
          <div className="content">
            {this.menu[this.state.selectedItem].component}
          </div>
        </div>
      </div>
    );
  }
}

export default App;
