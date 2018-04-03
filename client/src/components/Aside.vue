<template>
  <aside class="aside-menu">
    <b-tabs>
      <b-tab title="<i class='icon-list'></i>">
        <hr class="mx-3 my-0">
      </b-tab>
      <b-tab title="<i class='icon-speech'></i>">
        <div class="p-3">
          <div class="message">
            <div class="py-3 pb-5 mr-3 float-left">
              <div class="avatar">
                <img src="/static/img/avatars/7.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                <span class="avatar-status badge-success"></span>
              </div>
            </div>
            <div>
              <small class="text-muted">Lukasz Holeczek</small>
              <small class="text-muted float-right mt-1">1:52 PM</small>
            </div>
            <div class="text-truncate font-weight-bold">Lorem ipsum dolor sit amet</div>
            <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt...</small>
          </div>
          <hr>
          <div class="message">
            <div class="py-3 pb-5 mr-3 float-left">
              <div class="avatar">
                <img src="/static/img/avatars/7.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                <span class="avatar-status badge-success"></span>
              </div>
            </div>
            <div>
              <small class="text-muted">Lukasz Holeczek</small>
              <small class="text-muted float-right mt-1">1:52 PM</small>
            </div>
            <div class="text-truncate font-weight-bold">Lorem ipsum dolor sit amet</div>
            <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt...</small>
          </div>
          <hr>
          <div class="message">
            <div class="py-3 pb-5 mr-3 float-left">
              <div class="avatar">
                <img src="/static/img/avatars/7.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                <span class="avatar-status badge-success"></span>
              </div>
            </div>
            <div>
              <small class="text-muted">Lukasz Holeczek</small>
              <small class="text-muted float-right mt-1">1:52 PM</small>
            </div>
            <div class="text-truncate font-weight-bold">Lorem ipsum dolor sit amet</div>
            <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt...</small>
          </div>
          <hr>
          <div class="message">
            <div class="py-3 pb-5 mr-3 float-left">
              <div class="avatar">
                <img src="/static/img/avatars/7.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                <span class="avatar-status badge-success"></span>
              </div>
            </div>
            <div>
              <small class="text-muted">Lukasz Holeczek</small>
              <small class="text-muted float-right mt-1">1:52 PM</small>
            </div>
            <div class="text-truncate font-weight-bold">Lorem ipsum dolor sit amet</div>
            <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt...</small>
          </div>
          <hr>
          <div class="message">
            <div class="py-3 pb-5 mr-3 float-left">
              <div class="avatar">
                <img src="/static/img/avatars/7.jpg" class="img-avatar" alt="admin@bootstrapmaster.com">
                <span class="avatar-status badge-success"></span>
              </div>
            </div>
            <div>
              <small class="text-muted">Lukasz Holeczek</small>
              <small class="text-muted float-right mt-1">1:52 PM</small>
            </div>
            <div class="text-truncate font-weight-bold">Lorem ipsum dolor sit amet</div>
            <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt...</small>
          </div>
        </div>
      </b-tab>
      <b-tab title="<i class='icon-settings'></i>">
        <div class="p-3">
          <h6>Settings</h6>
          <div class="aside-options">
            <div class="clearfix mt-4">
              <small><b>Option 1</b></small>
              <c-switch type="text" variant="success" on="on" off="off" :pill="true" size="sm" class="float-right" :checked="true"/>
            </div>
            <div>
              <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</small>
            </div>
          </div>
          <div class="aside-options">
            <div class="clearfix mt-3">
              <small><b>Option 2</b></small>
              <c-switch type="text" variant="success" on="on" off="off" :pill="true" size="sm" class="float-right"/>
            </div>
            <div>
              <small class="text-muted">Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</small>
            </div>
          </div>
          <div class="aside-options">
            <div class="clearfix mt-3">
              <small><b>Option 3</b></small>
              <c-switch type="text" variant="success" on="on" off="off" :pill="true" size="sm" class="float-right"/>
            </div>
          </div>
          <div class="aside-options">
            <div class="clearfix mt-3">
              <small><b>Option 4</b></small>
              <c-switch type="text" variant="success" on="on" off="off" :pill="true" size="sm" class="float-right" :checked="true"/>
            </div>
          </div>
          <hr>
          <h6>System Utilization</h6>
          <div class="text-uppercase mb-1 mt-4"><small><b>CPU Usage</b></small></div>
          <b-progress class="progress-xs" variant="info" :value="25"></b-progress>
          <small class="text-muted">348 Processes. 1/4 Cores.</small>
          <div class="text-uppercase mb-1 mt-2"><small><b>Memory Usage</b></small></div>
          <b-progress class="progress-xs" variant="warning" :value="70"></b-progress>
          <small class="text-muted">11444GB/16384MB</small>
          <div class="text-uppercase mb-1 mt-2"><small><b>SSD 1 Usage</b></small></div>
          <b-progress class="progress-xs" variant="danger" :value="95"></b-progress>
          <small class="text-muted">243GB/256GB</small>
          <div class="text-uppercase mb-1 mt-2"><small><b>SSD 2 Usage</b></small></div>
          <b-progress class="progress-xs" variant="success" :value="10"></b-progress>
          <small class="text-muted">25GB/256GB</small>
        </div>
      </b-tab>
    </b-tabs>
  </aside>
</template>

<script>
import cSwitch from './Switch'
export default {
    name: 'app-aside',
    components: {
        cSwitch
    }
}
</script>
