import React, { useState } from "react";
import Brand from "./Brand";
import DisplayColumn from "./DisplayColumn";
import ButtonSpread from "./ButtonSpread";
import ModalManager from "./ModalManager";
import Layout from "./Layout";

const Display = props => {
  const [currentItem, setCurrentItem] = useState(
    props.items ? props.items[0] : null
  );

  return (
    <Layout>
      <div
        style={{
          width: "auto",
          height: "100vh",
          background:
            "linear-gradient(90deg, rgba(193,58,180,0.5956933798910189) 0%, rgba(253,29,29,0.6040967412355567) 50%, rgba(252,176,69,0.7525561249890581) 100%)"
        }}
      >
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
                  setCurrentItem={setCurrentItem}
                  onShowPreview={props.onShowPreview}
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
              header: "Question Preview",
              curr: currentItem
            }}
            question={{
              show: props.showQuestion,
              onHide: props.onHideQuestion,
              header: "Create Question"
            }}
            test={{
              show: props.showTest,
              onHide: props.onHideTest,
              header: "Create Test"
            }}
          />
        </div>
      </div>
    </Layout>
  );
};

export default Display;
