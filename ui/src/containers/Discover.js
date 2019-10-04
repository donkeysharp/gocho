import React, {Component} from 'react';
import Panel from '../components/Panel';
import NodeList from '../components/NodeList';
import NodeDetails from '../components/NodeDetails';

const refreshInterval = 3000;

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
    // Refresh the data every 3 secs
    this.refreshData = setInterval(this.retrieveData(), refreshInterval)
  }
  componentWillUnmount() {
    clearInterval(this.refreshData)
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
