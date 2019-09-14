import React, { useState } from "react";
import Brand from "./Brand";
import DisplayColumn from "./DisplayColumn";
import ButtonSpread from "./ButtonSpread";
import ModalManager from "./ModalManager";
import Layout from "./Layout";

const Display = props => {
  return (
    <Layout>
      <div className="p-4">
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
              <ButtonSpread buttons={props.buttons} />
            </div>
          </div>
        </div>
        <ModalManager
          question={{
            show: props.show,
            onHide: props.onHide,
            header: "Update Question"
          }}
        />
      </div>
    </Layout>
  );
};

export default Display;
