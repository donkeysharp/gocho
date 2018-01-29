import React, {Component} from 'react';

class NodeList extends Component {
  constructor(props) {
    super(props)
    this.state = {
      currentItem: -1
    }
  }
  onClickHandler(e) {
    let index = parseInt(e.currentTarget.dataset.index, 10);
    if (index === this.state.currentItem) {
      return;
    }
    this.setState({
      currentItem: index,
    })
    if (this.props.onNodeSelected) {
      this.props.onNodeSelected(index);
    }
  }
  render() {
    if (this.props.nodes && this.props.nodes.length === 0) {
      return <span>No nodes available</span>
    }
    return <ul className="node-list">
      {this.props.nodes.map((item, index) => {
        let className = ''
        if (index === this.state.currentItem) {
          className = 'active'
        }
        return <li key={index}>
          <a
            data-index={index}
            className={className}
            href="javascript:void(0)"
            onClick={this.onClickHandler.bind(this)}>
            {item.nodeId}
          </a>
        </li>
      })}
    </ul>
  }
}

export default NodeList;
