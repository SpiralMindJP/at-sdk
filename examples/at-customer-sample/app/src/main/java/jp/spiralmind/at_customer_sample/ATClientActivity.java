package jp.spiralmind.at_customer_sample;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.FrameLayout;

import jp.spiralmind.avatar_teleporter.customer.ATCustomerClientActivity;

public class ATClientActivity extends ATCustomerClientActivity{

    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        this.addControlsToUnityFrame();
        Intent intent = this.getIntent();
        this.handleIntent(intent);
    }

    @Override
    protected void onNewIntent(Intent intent) {
        super.onNewIntent(intent);
        this.handleIntent(intent);
        this.setIntent(intent);
    }

    private void handleIntent(Intent intent) {
        if (intent == null || intent.getExtras() == null) {
            return;
        }

        if (intent.getExtras().containsKey("doQuit")) {
            if (this.mUnityPlayer != null) {
                this.finish();
            }
        }
    }

    protected void showMainActivity() {
        Intent intent = new Intent(this, MainActivity.class);
        intent.setFlags(Intent.FLAG_ACTIVITY_REORDER_TO_FRONT | Intent.FLAG_ACTIVITY_SINGLE_TOP);
        this.startActivity(intent);
    }

    @Override
    public void onUnityPlayerUnloaded()
    {
        this.showMainActivity();
    }

    private void addControlsToUnityFrame() {
        FrameLayout layout = this.mUnityPlayer;

        {
            Button button = new Button(this);
            button.setText("Show Main");
            button.setX(10);
            button.setY(500);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    showMainActivity();
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Unload");
            button.setX(320);
            button.setY(500);
            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    mUnityPlayer.unload();
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Finish");
            button.setX(10);
            button.setY(560);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    finish();
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Set Viewport");
            button.setX(10);
            button.setY(620);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    SetViewport(0, 0, 1, 0.5f);
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Reset Viewport");
            button.setX(320);
            button.setY(620);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    SetViewport(0, 0, 1, 1);
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Show Setting View");
            button.setX(10);
            button.setY(680);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    ShowSettingView();
                }
            });
            layout.addView(button, 300, 50);
        }

        {
            Button button = new Button(this);
            button.setText("Show Select Server View");
            button.setX(320);
            button.setY(680);

            button.setOnClickListener(new View.OnClickListener() {
                public void onClick(View v) {
                    ShowSelectServerView();
                }
            });
            layout.addView(button, 300, 50);
        }

    }
}
