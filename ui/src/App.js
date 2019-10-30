import React, { Component } from 'react';
import { withTranslation } from 'react-i18next';
import SideBar from './components/SideBar';
import Discover from './containers/Discover';
import NodeInfo from './containers/NodeInfo';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.menu = [
      {
        name: 'menus.0.node_information',
        component: <NodeInfo />
      },
      {
        name: 'menus.1.discover',
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
    const { t } = this.props;
    return (
      <div className="wrapper">
        <span>{t("node_information")}</span>
        <SideBar
          toggle={this.state.toggle}
          title="Gocho"
          menu={this.menu}
          onMenuSelected={this.menuSelectedHandler.bind(this)} />
        <div className="content-wrapper">
          <nav className="nav">
            <b>{t(this.state.title)}</b>
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

export default withTranslation()(App);
