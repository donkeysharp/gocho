import React, {Component} from 'react';
import Panel from '../components/Panel';
import NodeList from '../components/NodeList';
import NodeDetails from '../components/NodeDetails';


class Discover extends Component {
  constructor(props) {
    super(props)
    this.state = {
      nodes: [],
      currentNode: -1,
    }
  }
  retrieveData() {
    fetch('/api/nodes').then((resp) => {
      return resp.json()
    }).then((data) => {
      this.setState({
        nodes: data
      })
    })
  }
  componentDidMount() {
    this.retrieveData()
  }
  nodeSelectedHandler(index) {
    this.setState({
      currentNode: index,
    });
  }
  render() {
    let detailsBody = <span>No node selected</span>
    if (this.state.currentNode !== -1) {
      detailsBody = <NodeDetails
        node={this.state.nodes[this.state.currentNode]}
      />
    }
    return <Panel title="Auto-Discovery">
      <div className="row">
        <div className="col-md-3">
          <NodeList
            nodes={this.state.nodes}
            onNodeSelected={this.nodeSelectedHandler.bind(this)} />
        </div>
        <div className="col-md-9">
          {detailsBody}
        </div>
      </div>
    </Panel>
  }
}

export default Discover;
