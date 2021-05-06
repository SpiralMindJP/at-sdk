package jp.spiralmind.at_customer_sample;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {
    private boolean isUnityLoaded = false;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }

    @Override
    protected void onNewIntent(Intent intent) {
        super.onNewIntent(intent);
        this.setIntent(intent);
    }

    public void onLoadButtonClick(View v) {
        this.isUnityLoaded = true;

        Intent intent = new Intent(this, ATClientActivity.class);
        intent.setFlags(Intent.FLAG_ACTIVITY_REORDER_TO_FRONT);
        this.startActivityForResult(intent, 1);
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);

        if (requestCode == 1) {
            this.isUnityLoaded = false;
        }
    }

    public void unloadUnity(Boolean doShowToast) {
        if(this.isUnityLoaded) {
            Intent intent = new Intent(this, ATClientActivity.class);
            intent.setFlags(Intent.FLAG_ACTIVITY_REORDER_TO_FRONT);
            intent.putExtra("doQuit", true);
            this.startActivity(intent);
            this.isUnityLoaded = false;
        }
        else if(doShowToast) {
            this.showToast("Show Unity First");
        }
    }

    public void onUnloadButtonClick(View v) {
        this.unloadUnity(true);
    }

    public void showToast(String message) {
        CharSequence text = message;
        int duration = Toast.LENGTH_SHORT;
        Toast toast = Toast.makeText(getApplicationContext(), text, duration);
        toast.show();
    }

    @Override
    public void onBackPressed() {
        this.finishAffinity();
    }
}