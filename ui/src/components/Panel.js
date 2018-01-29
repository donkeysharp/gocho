import React, {Component} from 'react';

class Panel extends Component {
  render() {
    return  <div className="panel">
      <header>
        {this.props.title}
      </header>
      <div className="panel-body">
        {this.props.children}
      </div>
    </div>
  }
}

export default Panel;
