import { waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { render } from 'test/test-utils';
import { byRole } from 'testing-library-selector';

import { setupDataSources } from 'app/features/alerting/unified/testSetup/datasources';
import { AccessControlAction } from 'app/types/accessControl';

import { setupMswServer } from '../../mockApi';
import { grantUserPermissions } from '../../mocks';
import { AlertmanagerProvider } from '../../state/AlertmanagerContext';

import AlertmanagerConfig from './AlertmanagerConfig';
import {
  EXTERNAL_VANILLA_ALERTMANAGER_UID,
  PROVISIONED_MIMIR_ALERTMANAGER_UID,
  mockDataSources,
  setupVanillaAlertmanagerServer,
} from './mocks/server';

const renderConfiguration = (
  alertManagerSourceName: string,
  { onDismiss = jest.fn(), onSave = jest.fn(), onReset = jest.fn() }
) =>
  render(
    <AlertmanagerProvider accessType="instance">
      <AlertmanagerConfig
        alertmanagerName={alertManagerSourceName}
        onDismiss={onDismiss}
        onSave={onSave}
        onReset={onReset}
      />
    </AlertmanagerProvider>
  );

const ui = {
  resetButton: byRole('button', { name: /Reset/ }),
  resetConfirmButton: byRole('button', { name: /Yes, reset configuration/ }),
  saveButton: byRole('button', { name: /Save/ }),
  cancelButton: byRole('button', { name: /Cancel/ }),
};

describe('Alerting Settings', () => {
  setupMswServer();

  beforeEach(() => {
    grantUserPermissions([AccessControlAction.AlertingNotificationsRead, AccessControlAction.AlertingInstanceRead]);
  });

  it('should not be able to reset alertmanager config', async () => {
    const onReset = jest.fn();
    renderConfiguration('grafana', { onReset });

    expect(ui.resetButton.query()).not.toBeInTheDocument();
  });

  it('should be able to cancel', async () => {
    const onDismiss = jest.fn();
    renderConfiguration('grafana', { onDismiss });

    await userEvent.click(await ui.cancelButton.get());
    expect(onDismiss).toHaveBeenCalledTimes(1);
  });
});

describe('vanilla Alertmanager', () => {
  const server = setupMswServer();

  beforeEach(() => {
    setupVanillaAlertmanagerServer(server);
    setupDataSources(...Object.values(mockDataSources));
    grantUserPermissions([AccessControlAction.AlertingNotificationsRead, AccessControlAction.AlertingInstanceRead]);
  });

  afterAll(() => {
    jest.resetAllMocks();
  });

  it('should be read-only when using vanilla Prometheus Alertmanager', async () => {
    renderConfiguration(EXTERNAL_VANILLA_ALERTMANAGER_UID, {});

    expect(ui.cancelButton.get()).toBeInTheDocument();
    expect(ui.saveButton.query()).not.toBeInTheDocument();
    expect(ui.resetButton.query()).not.toBeInTheDocument();
  });

  it('should not be read-only when Mimir Alertmanager', async () => {
    renderConfiguration(PROVISIONED_MIMIR_ALERTMANAGER_UID, {});

    expect(ui.cancelButton.get()).toBeInTheDocument();
    expect(ui.saveButton.get()).toBeInTheDocument();
    expect(ui.resetButton.get()).toBeInTheDocument();
  });

  it('should be able to reset non-Grafana alertmanager config', async () => {
    const onReset = jest.fn();
    renderConfiguration(PROVISIONED_MIMIR_ALERTMANAGER_UID, { onReset });

    expect(ui.cancelButton.get()).toBeInTheDocument();
    expect(ui.saveButton.get()).toBeInTheDocument();
    expect(ui.resetButton.get()).toBeInTheDocument();

    await userEvent.click(ui.resetButton.get());

    await userEvent.click(ui.resetConfirmButton.get());

    await waitFor(() => expect(onReset).toHaveBeenCalled());
    expect(onReset).toHaveBeenLastCalledWith(PROVISIONED_MIMIR_ALERTMANAGER_UID);
  });
});
