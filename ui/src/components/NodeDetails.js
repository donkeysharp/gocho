import React, {Component} from 'react';
import { withTranslation } from 'react-i18next';
import FormField from './FormField';


const IframeViewer = ((props) => {
  return <div className="row">
    <div className="col-md-12">
      <iframe frameBorder="0" className="files" src={props.url} title="shared-files" />
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
    const { t } = this.props;
    let nodeUrl = `http://${this.props.node.ipAddress}:${this.props.node.webPort}`;
    let iframeViewer = '';
    if (this.state.displayIframe) {
      iframeViewer = <IframeViewer url={nodeUrl} />
    }
    return <div className="node-details">
      <h4>{t('sections.discover.node_details')}</h4>
      <div className="form">
        <FormField label={t('sections.discover.node_id')} 
          value={this.props.node.nodeId}
          leftCol="col-md-2" rightCol="col-md-5" />
        <FormField label={t('sections.discover.web_port')} 
          value={this.props.node.webPort}
          leftCol="col-md-2" rightCol="col-md-5" />
        <FormField label={t('sections.discover.URL')} 
          value={nodeUrl}
          leftCol="col-md-2" rightCol="col-md-5" />
      </div>
      <div className="row">
        <div className="col-md-12">
          &nbsp;|&nbsp;
          <a
            href="#/"
            onClick={this.onClickHandler.bind(this)} >
            {
              this.state.displayIframe ? 
              t('sections.discover.hide_files') : 
              t('sections.discover.view_files')
            }
          </a>
          &nbsp;|&nbsp;
          <a href={nodeUrl} target="_">
            { t('sections.discover.open_in_tab') }
          </a>
          &nbsp;|&nbsp;
        </div>
      </div>
      { iframeViewer }
    </div>
  }
}

export default withTranslation()(NodeDetails);
