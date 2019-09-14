import React from "react";

const ButtonSpread = props => {
  return (
    <div className="text-center">
      {props.buttons.map((item, idx) => {
        return (
          <div className="row my-2" key={idx}>
            <div className="col-12">{item}</div>
          </div>
        );
      })}
    </div>
  );
};

export default ButtonSpread;
