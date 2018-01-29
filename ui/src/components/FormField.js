import React, {Component} from 'react';

class FormField extends Component {
  render() {
    let disabled = 'false';
    if (this.props.isDisabled) {
      disabled = 'true';
    }
    let leftCol = this.props.leftCol ? this.props.leftCol : 'col-md-3';
    let rightCol = this.props.rightCol ? this.props.rightCol : 'col-md-9';
    return <div className="form-group row">
      <label className={leftCol + ' col-form-label'}>{this.props.label}</label>
      <div className={rightCol}>
        <input className="form-control" disabled={disabled} value={this.props.value} />
      </div>
    </div>;
  }
}

export default FormField;
