import React, {Component} from 'react';
import FormField from './FormField';


const IframeViewer = ((props) => {
  return <div className="row">
    <div className="col-md-12">
      <iframe frameBorder="0" className="files" src={props.url} />
    </div>
  </div>
})

class NodeDetails extends Component {
  constructor(props) {
    super(props)
    this.state = {
      displayIframe: false,
    }
  }
  componentWillReceiveProps(nextProps) {
    if (this.props.node.nodeId !== nextProps.node.nodeId) {
      this.setState({
        displayIframe: false,
      })
    }
  }
  onClickHandler() {
    this.setState({
      displayIframe: !this.state.displayIframe,
    })
  }
  render() {
    let nodeUrl = `http://${this.props.node.ipAddress}:${this.props.node.webPort}`;
    let iframeViewer = '';
    if (this.state.displayIframe) {
      iframeViewer = <IframeViewer url={nodeUrl} />
    }
    return <div className="node-details">
      <h4>Node Details</h4>
      <div className="form">
        <FormField label="Node ID" value={this.props.node.nodeId}
          leftCol="col-md-2" rightCol="col-md-5" />
        <FormField label="Web Port" value={this.props.node.webPort}
          leftCol="col-md-2" rightCol="col-md-5" />
        <FormField label="URL" value={nodeUrl}
          leftCol="col-md-2" rightCol="col-md-5" />
      </div>
      <div className="row">
        <div className="col-md-12">
          &nbsp;|&nbsp;
          <a
            href="javascript:void(0)"
            onClick={this.onClickHandler.bind(this)} >
            {this.state.displayIframe ? 'Hide files' : 'View files'}
          </a>
          &nbsp;|&nbsp;
          <a href={nodeUrl} target="_">Open in tab</a>
          &nbsp;|&nbsp;
        </div>
      </div>
      { iframeViewer }
    </div>
  }
}

export default NodeDetails;
