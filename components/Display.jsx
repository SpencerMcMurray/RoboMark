import React from "react";
import Brand from "./Brand";
import DisplayColumn from "./DisplayColumn";
import ButtonSpread from "./ButtonSpread";
import Layout from "./Layout";

const DisplayLayout = props => {
  return (
    <Layout>
      <div className="d-flex justify-content-center">
        <div className="row w-100" style={{ height: "90vh" }}>
          <div className="col-lg-3 text-center">
            <Brand />
          </div>
          <div className="col-lg-6 d-flex align-items-center">
            <DisplayColumn
              title={props.title}
              headers={props.headers}
              items={props.items}
            />
          </div>
          <div className="col-lg-3">
            <ButtonSpread />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default DisplayLayout;
