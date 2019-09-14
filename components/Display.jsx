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
            <div className="col-lg-3 d-flex align-items-center justify-content-center">
              <ButtonSpread buttons={props.buttons} />
            </div>
          </div>
        </div>
        <ModalManager
          preview={{
            show: props.showPreview,
            onHide: props.onHidePreview,
            header: "Question Preview"
          }}
          question={{
            show: props.showQuestion,
            onHide: props.onHideQuestion,
            header: "Create Question"
          }}
          page={{
            show: props.showPage,
            onHide: props.onHidePage,
            header: "Create Page"
          }}
          test={{
            show: props.showTest,
            onHide: props.onHideTest,
            header: "Create Test"
          }}
        />
      </div>
    </Layout>
  );
};

export default Display;
